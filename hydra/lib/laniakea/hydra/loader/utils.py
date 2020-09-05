import importlib
import io
import types
import zipfile
from abc import ABC, abstractmethod
from importlib.machinery import ModuleSpec
from typing import Optional, Sequence, Union

from . import logger


class ResourceLoader(ABC):

    @abstractmethod
    def get_name(self) -> str:
        pass

    @abstractmethod
    def get_resource(self, name: str) -> bytes:
        pass

    @abstractmethod
    def has_resource(self, name: str) -> bool:
        pass


class Loader(ABC):

    @abstractmethod
    def get_base(self) -> str:
        pass

    @abstractmethod
    def get_resource_loader(self, name: str) -> Optional[ResourceLoader]:
        pass


class ZipFileResourceLoader(ResourceLoader):
    def __init__(self, name: str, contents: bytes):
        buf = io.BytesIO(contents)
        self.zip_file = zipfile.ZipFile(buf)
        self.name = name

    def get_name(self) -> str:
        return self.name

    def get_resource(self, name: str) -> bytes:
        return self.zip_file.read(name)

    def has_resource(self, name: str) -> bool:
        try:
            self.zip_file.getinfo(name)
            return True
        except KeyError:
            return False


class Importer(importlib.abc.MetaPathFinder, importlib.abc.Loader):
    def __init__(self, loader: Loader):
        self.loader = loader

    def find_spec(self, fullname: str, path: Optional[Sequence[Union[bytes, str]]],
                  target: Optional[types.ModuleType] = ...) -> \
            Optional[ModuleSpec]:
        if fullname.startswith("laniakea.hydra.dyn"):
            logger.debug(f"Loading {fullname}")
            name_parts = fullname.split('.')
            if len(name_parts) == 3:
                return ModuleSpec(name=fullname, loader=self,
                                  origin=f"{self.loader.get_base()}", is_package=True)
            elif len(name_parts) == 4:
                dist_name = fullname.split('.')[3]
                rl = self.loader.get_resource_loader(dist_name)
                if rl:
                    return ModuleSpec(name=fullname, loader=self,
                                      origin=f"{self.loader.get_base()}{rl.get_name()}", is_package=True)
                else:
                    return None
            else:
                dist_name = fullname.split('.')[3]
                package_path = '/'.join(fullname.split('.')[4:])
                rl = self.loader.get_resource_loader(dist_name)
                if rl:
                    if rl.has_resource(package_path + ".py"):
                        return ModuleSpec(name=fullname, loader=self,
                                          origin=f"{self.loader.get_base()}{rl.get_name()}", is_package=False)
                    elif rl.has_resource(package_path + "/__init__.py"):
                        return ModuleSpec(name=fullname, loader=self,
                                          origin=f"{self.loader.get_base()}{rl.get_name()}", is_package=True)
                    else:
                        return None
        else:
            return None

    def exec_module(self, module: types.ModuleType) -> None:
        fullname = module.__name__
        logger.debug(f"Executing {fullname}")
        name_parts = fullname.split('.')
        if len(name_parts) > 3:
            package_path = '/'.join(fullname.split('.')[4:])
            rl = self.loader.get_resource_loader(name_parts[3])
            if rl.has_resource(package_path + ".py"):
                mod_source = rl.get_resource(package_path + ".py")
            elif rl.has_resource(package_path + "/__init__.py"):
                print("loading source:" + package_path + "/__init__.py")
                mod_source = rl.get_resource(package_path + "/__init__.py")
            else:
                return None
            exec(mod_source, module.__dict__)
        return None
