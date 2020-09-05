import sys
from functools import lru_cache
from typing import Optional

from google.cloud import storage

from .utils import Loader, Importer, ResourceLoader, ZipFileResourceLoader


def add_gcs(bucket: str, path: str, project: str = None):
    loader = GCSLoader(bucket, path, project)
    importer = Importer(loader)
    sys.meta_path.insert(0, importer)


class GCSLoader(Loader):
    def __init__(self, bucket: str, path: str, project: str = None):
        if project:
            self.bucket = storage.Client(project=project).bucket(bucket)
        else:
            self.bucket = storage.Client().bucket(bucket)
        self.path = path.rstrip('/').lstrip('/')+'/'
        self.base = f"gs://{self.bucket.name}/{self.path}"

    @lru_cache(maxsize=64)
    def get_resource_loader(self, name: str) -> Optional[ResourceLoader]:
        blob = self.bucket.get_blob(self.path + name + '.zip')
        if blob:
            return ZipFileResourceLoader(name=name + '.zip', contents=blob.download_as_string())
        else:
            return None

    def get_base(self) -> str:
        return self.base
