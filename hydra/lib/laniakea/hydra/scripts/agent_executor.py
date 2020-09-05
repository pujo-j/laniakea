import prefect
from prefect.client import Client
from prefect.engine import get_default_executor_class
from prefect.environments import LocalEnvironment
from prefect.tasks.secrets import PrefectSecret
from prefect.utilities.graphql import with_args


def execute():
    flow_run_id = prefect.context.get("flow_run_id")
    if not flow_run_id:
        print("Not currently executing a flow within a Cloud context.")
        raise Exception("Not currently executing a flow within a Cloud context.")

    query = {
        "query": {
            with_args("flow_run", {"where": {"id": {"_eq": flow_run_id}}}): {
                "flow": {"name": True, "storage": True, "environment": True},
                "version": True,
            }
        }
    }

    client = Client()
    result = client.graphql(query)
    flow_run = result.data.flow_run

    if not flow_run:
        print("Flow run {} not found".format(flow_run_id))
        raise ValueError("Flow run {} not found".format(flow_run_id))

    try:
        flow_data = flow_run[0].flow
        storage_schema = prefect.serialization.storage.StorageSchema()
        storage = storage_schema.load(flow_data.storage)

        # populate global secrets
        secrets = prefect.context.get("secrets", {})
        for secret in storage.secrets:
            secrets[secret] = PrefectSecret(name=secret).run()

        with prefect.context(secrets=secrets, loading_flow=True):
            flow = storage.get_flow(storage.flows[flow_data.name])
            environment = LocalEnvironment()
            environment.executor = get_default_executor_class()()

            environment.setup(flow)
            environment.execute(flow)
    except Exception as exc:
        msg = "Failed to load and execute Flow's environment: {}".format(repr(exc))
        state = prefect.engine.state.Failed(message=msg)
        client.set_flow_run_state(flow_run_id=flow_run_id, state=state)
        print(str(exc))
        raise exc
