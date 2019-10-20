#!/usr/bin/env python3
import yaml
import sys
from jinja2 import Environment, FileSystemLoader

node_id = sys.argv[1]
if __name__ == "__main__":
 # yaml = YAML()
 config_data = yaml.load(open('./envoy.yaml.j2'))
 print(config_data)
# Load templates file from templtes folder
 env = Environment(loader = FileSystemLoader('./templates'), trim_blocks=True, lstrip_blocks=True)
 template = env.get_template('template.txt')
 print(template.render(config_data))