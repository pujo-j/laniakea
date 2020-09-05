from distributed import WorkerPlugin

import laniakea.hydra.loader


class LoaderPlugin(WorkerPlugin):
    def __init__(self):
        pass

    def setup(self, worker):
        laniakea.hydra.loader.init()
