import laniakea.hydra.loader

laniakea.hydra.loader.init()

import base64
import datetime
import hmac
import logging
import os
import struct
from typing import Dict

import click
import git
import pendulum
import prefect
from prefect.environments.storage import Webhook
from prefect.utilities.storage import extract_flow_from_file
from slugify import slugify

from laniakea.hydra.lineage.handler import get_lineage_handler_for_sink
from laniakea.hydra.lineage.sink_discovery import DynamicLineageSink
from laniakea.hydra.scripts import init_logging
from laniakea.hydra.scripts.agent import LnkAgent
from laniakea.hydra.scripts.agent_executor import execute as agent_execute

logger = logging.getLogger("cli")


@click.group()
def main():
    pass


@main.command()
@click.option(
    "--file",
    "-f",
    required=True,
    help="A file that contains a flow",
    type=click.Path(exists=True),
)
@click.option(
    "--project",
    "-p",
    required=True,
    help="The name of a Prefect project to register this flow.",
)
def register(file, project):
    init_logging()
    with prefect.context({"loading_flow": True}):
        file_path = os.path.abspath(file)
        flow = extract_flow_from_file(file_path=file_path)
    try:
        repo = git.Repo(search_parent_directories=True)
        version = repo.head.object.hexsha
    except git.exc.InvalidGitRepositoryError:
        version = slugify(pendulum.now('utc').isoformat())
    for task in flow.get_tasks():
        if task.slug:
            task.slug = task.slug.rstrip("copy") + version
        else:
            task.slug = slugify(task.name) + "-" + version
        if task.state_handlers:
            task.state_handlers = task.state_handlers.append(get_lineage_handler_for_sink(DynamicLineageSink()))
        else:
            task.state_handlers = [task.state_handlers.append(get_lineage_handler_for_sink(DynamicLineageSink()))]
    key = f"{slugify(flow.name)}{version}"
    w = Webhook(
        stored_as_script=True,
        flow_script_path=file_path,
        build_request_kwargs={
            "url": "${LANIAKEA_STORAGE}" + key,
            "headers": {
                "Authorization": "Bearer ${PREFECT__CLOUD__AUTH_TOKEN}",
            }
        },
        build_request_http_method="PUT",
        get_flow_request_kwargs={
            "url": "${LANIAKEA_STORAGE}" + key,
            "headers": {
                "Authorization": "Bearer ${PREFECT__CLOUD__AUTH_TOKEN}",
            }
        },
        get_flow_request_http_method="GET"
    )

    flow.storage = w
    flow.register(project_name=project, no_url=True)


@main.command()
@click.option('--key',
              help="Server secret key",
              prompt=True)
@click.option('--user',
              help="User/Service name",
              prompt=True)
@click.option('--expire',
              help="Days before token expiry",
              default=365)
def token(key, user, expire):
    expiry = datetime.datetime.utcnow() + datetime.timedelta(days=expire)
    exp_buf = struct.pack('<Q', int(expiry.timestamp()))
    to_sign = exp_buf + user.encode('UTF-8')
    key_bytes = base64.standard_b64decode(key + '==')
    sign = hmac.new(key_bytes, digestmod='SHA256', msg=to_sign).digest()
    token_bytes = exp_buf + sign + user.encode('UTF-8')
    token_string = base64.standard_b64encode(token_bytes).rstrip(b'=')
    click.echo(token_string)


@main.command(hidden=True)
def agent():
    init_logging()
    logger.info("Starting agent")
    import sys
    print("ARGGG AGENT")
    print(f"{sys.meta_path=}")
    if os.getenv("DEBUG"):
        LnkAgent(show_flow_logs=True).start()
    else:
        LnkAgent().start()


@main.command(hidden=True)
def execute():
    init_logging()
    logger.info("Starting executor")
    agent_execute()


def extract_flow(entities: Dict) -> "prefect.Flow":
    for var in entities:
        if isinstance(entities[var], prefect.Flow):
            return entities[var]

    raise ValueError("No flow found in file.")


def extract_entities_from_file(file_path: str) -> Dict:
    if file_path:
        with open(file_path, "r") as f:
            contents = f.read()
    file_entities = {}
    exec(contents, file_entities)
    return file_entities
