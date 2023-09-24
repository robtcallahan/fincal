variable "region" {
  type        = string
  description = "AWS Region for S3 bucket"
}

variable "availability_zones" {
  type        = list(string)
  description = "List of availability zones"
}

variable "vpc_security_group_ids" {
  type        = list(string)
  description = "List of security groups"
}

variable "subnet_id" {
  type        = string
  description = "AWS subnet id"
}

variable "instance_name" {
  description = "Value of the Name tag for the EC2 instance"
  type        = string
}
