# OUT_DIR: Temporary folder for image data
# VERSION: Laniakea Version
# BUCKET: Base GCS bucket for artifact storage

export IMG_PREFIX="laniakea"
mkdir -p $OUT_DIR/$VERSION

for IMAGE in ui storage apollo server wait; do
  pushd $IMAGE
  docker build -t $IMG_PREFIX/$IMAGE:$VERSION -f Dockerfile .
  popd
  docker save $IMG_PREFIX/$IMAGE:$VERSION | pbzip2 -c >$OUT_DIR/$VERSION/$IMAGE.tar.bz2
done

docker pull postgres:12.3-alpine
docker tag postgres:12.3-alpine $IMG_PREFIX/sql:$VERSION
docker pull hasura/graphql-engine:v1.1.0
docker tag hasura/graphql-engine:v1.1.0 $IMG_PREFIX/hasura:$VERSION
docker pull busybox:latest

docker save $IMG_PREFIX/sql:$VERSION | pbzip2 -c >$OUT_DIR/$VERSION/sql.tar.bz2
docker save $IMG_PREFIX/hasura:$VERSION | pbzip2 -c >$OUT_DIR/$VERSION/hasura.tar.bz2
docker save busybox:latest | pbzip2 -c >$OUT_DIR/$VERSION/busybox.tar.bz2

gsutil -m rsync -r -d $OUT_DIR/$VERSION/ $BUCKET/lnk-norma-$VERSION/

helm package ../norma -d $OUT_DIR/$VERSION/

gsutil -m copy $OUT_DIR/$VERSION/norma-$VERSION.tgz $BUCKET/lnk-norma-$VERSION/
