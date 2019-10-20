#!/usr/bin/env python3
import yaml
import sys
import os
from jinja2 import Environment, FileSystemLoader

basepath = os.path.abspath(os.path.dirname(__file__))
config_data = yaml.load(open(basepath + '/data.txt'))
# Load templates file from templtes folder
env = Environment(loader = FileSystemLoader(basepath), trim_blocks=True, lstrip_blocks=True)
template = env.get_template('envoy.yaml.j2')
print(template.render(config_data))