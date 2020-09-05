# OUT_DIR: Temporary folder for image data
# VERSION: Laniakea Version
# BUCKET: Base GCS bucket for artifact storage

export IMG_PREFIX="laniakea"
mkdir -p $OUT_DIR/hydra/$VERSION

# Copy laniakea to image folder
pushd  ../lib/
poetry build
mkdir -p ../images/worker-cpu/laniakea/
cp -r . ../images/worker-cpu/laniakea/
gsutil -m copy dist/laniakea-$VERSION.tar.gz $BUCKET/lnk-hydra-$VERSION/
popd

pushd worker-cpu
docker build -t $IMG_PREFIX/worker-cpu:$VERSION -f Dockerfile .
popd
docker save $IMG_PREFIX/worker-cpu:$VERSION | pbzip2 -c >$OUT_DIR/hydra/$VERSION/worker-cpu.tar.bz2

docker pull busybox:latest

docker save busybox:latest | pbzip2 -c >$OUT_DIR/hydra/$VERSION/busybox.tar.bz2

gsutil -m rsync -r -d $OUT_DIR/hydra/$VERSION/ $BUCKET/lnk-hydra-$VERSION/

helm package ../hydra -d $OUT_DIR/hydra/$VERSION/

gsutil -m copy $OUT_DIR/hydra/$VERSION/hydra-$VERSION.tgz $BUCKET/lnk-hydra-$VERSION/
