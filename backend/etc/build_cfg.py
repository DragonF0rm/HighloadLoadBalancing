#!/usr/bin/env python3
import yaml
import sys
from jinja2 import Environment, FileSystemLoader

basepath = os.path.abspath(os.path.dirname(__file__))
config_data = yaml.load(open(basepath + '/envoy.yaml.j2'))
# Load templates file from templtes folder
env = Environment(loader = FileSystemLoader('./templates'), trim_blocks=True, lstrip_blocks=True)
template = env.get_template('template.txt')
print(template.render(config_data))