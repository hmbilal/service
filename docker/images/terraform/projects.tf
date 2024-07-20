resource "aws_dynamodb_table" "projects" {
  name           = "projects"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "AccessKey"

  attribute {
    name = "AccessKey"
    type = "S"
  }
}
