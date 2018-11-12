workflow AppStack {
  typespace => 'genesis::aws',
  input => (String $region = lookup('aws::region', Hash[String,String] $tags = lookup('aws::tags')),
  output => ($vpc_id, $subnet_id)
} {
  resource vpc {
    output => ($vpc_id)
  } {
    region               => $region,
    cidr_block           => '192.168.0.0/16',
    tags                 => $tags,
    enable_dns_hostnames => true,
    enable_dns_support   => true
  }

  resource subnet {
    output => ($subnet_id)
  } {
    region               => $region,
    vpc_id               => $vpc_id,
    cidr_block           => '192.168.1.0/24',
    tags                 => $tags,
    map_public_ip_on_launch => true
  }
  
  resource securitygroup {
    output => ($securitygroup_id)
  } {
    region => $region,
    rules  => $rules,
    tags   => $tags
  }

  resource launch_configuration {
    output => ($launchconfig_id)
  } {
    name                        => "master-us-west-2a-masters-kenazcluster-k8s-local"
    name_prefix                 => "master-us-west-2a.masters.kenazcluster.k8s.local-"
    image_id                    => "ami-4bfe6f33"
    instance_type               => "m3.medium"
    key_name                    => "keypair.id"
    security_groups             => $securitygroup_id
    associate_public_ip_address => true
    root_block_device = {
      volume_type           => "gp2"
      volume_size           => 64
      delete_on_termination => true
    }
    ephemeral_block_device => {
      device_name  => "/dev/sdc"
      virtual_name => "ephemeral0"
    }
    lifecycle => {
      create_before_destroy => true
    }
    enable_monitoring => false
  }

  resource autoscaling_group {
    name                 => "nodes.kenazcluster.k8s.local"
    launch_configuration => $launchconfig_id
    max_size             => 3
    min_size             => 3
    vpc_zone_identifier  => $subnet_id
}
