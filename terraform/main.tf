terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = var.region
}

resource "aws_key_pair" "tf-key-pair" {
  key_name   = "tf-key-pair"
  public_key = tls_private_key.rsa.public_key_openssh
}
resource "tls_private_key" "rsa" {
  algorithm = "RSA"
  rsa_bits  = 4096
}
resource "local_file" "tf-key" {
  content  = tls_private_key.rsa.private_key_pem
  filename = "tf-key-pair"
}

resource "aws_instance" "app_server" {
  # default ami
#  ami                    = "ami-03a6eaae9938c858c"
  # fincal-container-host
  ami                    = "ami-0454d47ad07f6ff7f"
  instance_type          = "t2.micro"
  vpc_security_group_ids = var.vpc_security_group_ids
  subnet_id              = var.subnet_id
  key_name               = "tf-key-pair"

  tags = {
    Name = var.instance_name
  }
}

