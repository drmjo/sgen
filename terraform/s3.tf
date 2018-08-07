resource "aws_s3_bucket" "website" {
  provider = "aws.west"
  bucket = "${var.domain}"
  acl    = "private"
}

# resource "aws_s3_bucket_policy" "website_policy" {
#   	bucket = "${aws_s3_bucket.website.id}"
#     policy = "${data.aws_iam_policy_document.s3_access.json}"
# }
