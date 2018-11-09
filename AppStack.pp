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
}
