provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_s3_bucket" "bucket" {
  bucket = "saotomek-playground-bucket"
  acl    = "private"
  object_lock_configuration {
    object_lock_enabled = "Enabled"
    rule {
      default_retention {
        mode = "GOVERNANCE"
        days = 3650
      }
    }
  }
}
