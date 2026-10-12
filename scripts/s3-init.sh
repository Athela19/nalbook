#!/bin/sh
set -u
ENDPOINT="${S3_ENDPOINT:?S3_ENDPOINT wajib diisi}"
BUCKET="${S3_BUCKET:?S3_BUCKET wajib diisi}"
REGION="${S3_REGION:-us-east-1}"
ATTEMPTS="${S3_INIT_ATTEMPTS:-30}"
WAIT_SECONDS=2
n=1
while [ "$n" -le "$ATTEMPTS" ]; do
  if OUT=$(aws --endpoint-url "$ENDPOINT" --region "$REGION" s3api head-bucket --bucket "$BUCKET" 2>&1); then
    echo "bucket $BUCKET sudah ada"
    exit 0
  fi
  case "$OUT" in
    *404*|*NoSuchBucket*|*"Not Found"*)
      if MB_OUT=$(aws --endpoint-url "$ENDPOINT" --region "$REGION" s3 mb "s3://$BUCKET" 2>&1); then
        echo "bucket $BUCKET dibuat"
        exit 0
      fi
      case "$MB_OUT" in
        *BucketAlreadyOwnedByYou*|*BucketAlreadyExists*)
          echo "bucket $BUCKET sudah ada"
          exit 0
          ;;
      esac
      echo "$MB_OUT"
      exit 1
      ;;
    *"Could not connect"*|*"Connect timeout"*|*"Max retries exceeded"*|*"Name or service not known"*|*"Temporary failure in name resolution"*)
      echo "percobaan $n/$ATTEMPTS: endpoint belum siap"
      sleep "$WAIT_SECONDS"
      n=$((n + 1))
      ;;
    *)
      echo "$OUT"
      exit 1
      ;;
  esac
done
echo "endpoint $ENDPOINT tidak merespons setelah $ATTEMPTS percobaan"
exit 1
