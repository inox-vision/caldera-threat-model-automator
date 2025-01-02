#! /bin/python

import os
import yaml
import json


# Adversary profile name

print(f'''\n    [!] This script creates MITRE Caldera adversary profile, 
  based on chosen platforms and techniques, using srockpile
  and atomic plugins and JSON file generated with CVC:
  https://controlcompass.github.io/threat-model .\n''')
  
adv_name = input(f"    [>] Type in your Caldera adversary profile name:\n        ")
descr_choice = input(f"    [?] Do you want to add custom description? (y/n)\n        ")
if descr_choice in ('y', 'Y'):
    description = input(f"    [>] Custom description:\n        ")
else:
    description = ''



# JSON Threat Profile parsing

threat_file = input("Type in Adversary profile json file name: ")
json_techniques = []

with open(threat_file, "r") as json_file:
    json_stream = json.load(json_file)

for t in json_stream['techniques']:
    json_techniques.append(t['techniqueID'])


# File creation

adv_file = open(adv_name + '.yml', 'w')
adv_file.write(f"""---

name: {adv_name}
description: {description}
atomic_ordering:""")


# techniques

techniques = [ 
    'reconnaissance', 
    'initial-access', 
    'defense-evasion', 
    'persistence',
    'command-and-control',
    'discovery',
    'privilege-escalation',
    'execution',
    'credential-access',
    'lateral-movement',
    'collection',
    'exfiltration',
    'impact',
    'multiple',
    ]


# Platforms for which adversary profile should be created.

platforms = [
    'linux',
    'darwin',
    'windows',
    ]

platforms_chosen = ['unknown']

print('    [>] Choose platforms you want to create adversary profile for:')
for platform in platforms:
    choice = input(f'{platform} ? y/n\n')
    if choice in ('y', 'Y'):
        platforms_chosen.append(platform)


stockpile_directory = "./caldera/plugins/stockpile/data/abilities"
atomic_directory = "./caldera/plugins/atomic/data/abilities"

print('    [!] Adversary profile is being created.')

for technique in techniques:
    for dirpath, dirs, files in os.walk(stockpile_directory + '/' + technique):
        for file in files:
            file_path = os.path.join(dirpath, file)
            with open(file_path, "r") as ability: 
                    data = yaml.safe_load(ability)
            try:
                ab_platforms = data[0]['platforms'].keys()
                ID = data[0]['id']
                attack_ID = data[0]['technique']['attack_id']
                if any(item in ab_platforms for item in platforms_chosen):
                    if attack_ID in json_techniques:
                        adv_file = open(adv_name + '.yml', 'a')
                        adv_file.write(f'\n - {ID}')
                
            except KeyError:   
                pass             
                
    for dirpath, dirs, files in os.walk(atomic_directory + '/' + technique):
        for file in files:
            file_path = os.path.join(dirpath, file)
            with open(file_path, "r") as ability: 
                    data = yaml.safe_load(ability)
            try:
                ab_platforms = data[0]['platforms'].keys()
                ID = data[0]['id']
                attack_ID = data[0]['technique']['attack_id']

                
                if any(item in ab_platforms for item in platforms_chosen):
                    if attack_ID in json_techniques:
                        adv_file = open(adv_name + '.yml', 'a')
                        adv_file.write(f'\n - {ID}')
                
            except KeyError:   
                pass

print(f'    [!] Success! You can use your {adv_name} file in Caldera')