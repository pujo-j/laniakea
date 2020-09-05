import os
import sys
from typing import Optional

from .utils import Loader, Importer, ResourceLoader


def add_path(dir_path: str):
    loader = PathLoader(dir_path)
    importer = Importer(loader)
    sys.meta_path.insert(0, importer)


class DirResourceLoader(ResourceLoader):

    def __init__(self, name: str, base_path: str):
        self.base_path = base_path
        self.name = name

    def get_name(self) -> str:
        return self.name

    def get_resource(self, name: str) -> bytes:
        with open(os.path.join(self.base_path, name), "rb") as fd:
            return fd.read()

    def has_resource(self, name: str) -> bool:
        return os.path.exists(os.path.join(self.base_path, name))


class PathLoader(Loader):

    def __init__(self, dir_path: str):
        self.dir_path = os.path.abspath(dir_path) + '/'
        self.base = f"file://{self.dir_path}"

    def get_base(self) -> str:
        return self.base

    def get_resource_loader(self, name: str) -> Optional[ResourceLoader]:
        rp = os.path.join(self.dir_path, name)
        if os.path.exists(rp) and os.path.isdir(rp):
            return DirResourceLoader(name, rp)
        else:
            return None
