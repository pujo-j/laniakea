import logging

logger = logging.getLogger("laniakea.hydra.loader")

loaded = False


def init():
    global loaded
    if not loaded:
        import os

        from .fs import add_path
        from .gcs import add_gcs
        from .http import add_http
        repo = os.getenv("HYDRA_REPO")
        if repo and repo != "":
            logger.debug(f"Using hydra repos: {repo}")
            for r in repo.split(','):
                if r.startswith('http'):
                    proxy = os.getenv("http_proxy")
                    auth_token = os.getenv("HYDRA_REPO_TOKEN")
                    logger.info(f"Adding library HTTP repository:{r}")
                    if auth_token:
                        add_http(r, proxy, headers={
                            "Authorization": "Bearer " + auth_token,
                        })
                    else:
                        add_http(r, proxy)
                elif r.startswith('gs://'):
                    bucket = r[5:].split('/')[0]
                    path = '/'.join(r[5:].split('/')[1:])
                    logger.info(f"Adding library GCS repository:{r}")
                    add_gcs(bucket=bucket, path=path)
                elif r.startswith('file://'):
                    path = '/'.join(r[6:].split('/')[1:])
                    logger.info(f"Adding library File repository:{r}")
                    add_path(path)
                else:
                    # TODO: Add S3 and Azure Blob Storage
                    raise Exception(f"invalid HYDRA_REPO entry: {r}")
        loaded = True
