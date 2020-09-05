import sys
from functools import lru_cache
from typing import Optional, Dict

import requests

from .utils import Loader, Importer, ResourceLoader, ZipFileResourceLoader


def add_http(url: str, proxy: Optional[str] = None, headers: Optional[Dict[str, str]] = None):
    loader = HttpLoader(url=url, proxy=proxy, headers=headers)
    importer = Importer(loader)
    sys.meta_path.insert(0, importer)


class HttpLoader(Loader):
    def __init__(self, url: str, proxy: Optional[str] = None, headers: Optional[Dict[str, str]] = None):
        if not headers:
            self.headers = {}
        else:
            self.headers = headers
        if proxy and proxy != "":
            self.proxy = proxy
        else:
            self.proxy = None
        self.url = url
        self.base = f"{self.url}"

    @lru_cache(maxsize=64)
    def get_resource_loader(self, name: str) -> Optional[ResourceLoader]:
        if self.proxy:
            res = requests.get(self.url + name + ".zip", proxies={
                'http': self.proxy,
                'https': self.proxy,
            }, headers=self.headers)
        else:
            res = requests.get(self.url + name + ".zip", headers=self.headers)
        if res.status_code != 200:
            res.close()
            return None
        return ZipFileResourceLoader(name=name + '.zip', contents=res.content)

    def get_base(self) -> str:
        return self.base
