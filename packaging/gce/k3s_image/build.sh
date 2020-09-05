packer build \
  -var project_id=$PROJECT_ID \
  -var source_bucket=$BUCKET \
  -var sac=$SAC \
  -var image_zone=$ZONE \
  -var image_name=k3s-$(date +%Y%m%d%H%M) \
  packer.json
