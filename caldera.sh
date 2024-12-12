#!/bin/bash

# Initial input
echo "Type in your Caldera adversary model name:"
read adv_name 
echo "CVC json threat model file name:"
read -e model_file
echo "Specify platforms this adversary is created to test against:"
read -p "MacOS? [y/n]" darwin
read -p "Windows? [y/n]" windows
read -p "Linux? [y/n]" linux
touch platforms

# Creating platforms input for abilities parsing
if [ $darwin == "y" ]
    then
    echo "darwin" >> platforms
fi

if [ $windows == "y" ]
    then
    echo "windows" >> platforms
fi

if [ $linux == "y" ]
    then
    echo "linux" >> platforms
fi

# Parsing Caldera abilities files
echo "    [!] Parsing Caldera abilities" 

output=abilities_parsed
echo "" > $output

# Specyfying paths to Caldera plugins abilities
stockpile="./caldera/plugins/stockpile/data/abilities"
atomic="./caldera/plugins/atomic/data/abilities"

abilities () {
    for file in $(find $1/$2 -name "*.yml")
do
    if [[ $(cat $file | grep -f platforms) ]]
    then
        echo "$(cat $file | sed -n 's/.* id: \(.*\)/\1/p') $(cat $file | sed -n 's/.*attack_id: \(.*\)/\1/p') $(cat $file | sed -n 's/.*technique_id: \(.*\)/\1/p') $(cat $file | sed -n 's/.* tactic: \(.*\)/\1/p')" >> $output
    fi
done
}

# Order corresponds more or less to Unified Kill Chain
abilities $atomic reconnaissance
abilities $atomic initial-access
abilities $stockpile defense-evasion
abilities $atomic defense-evasion
abilities $stockpile persistence
abilities $atomic persistence
abilities $stockpile command-and-control 
abilities $atomic command-and-control
abilities $stockpile discovery
abilities $atomic discovery
abilities $stockpile privilege-escalation 
abilities $atomic privilege-escalation
abilities $stockpile execution 
abilities $atomic execution
abilities $stockpile credential-access 
abilities $atomic credential-access
abilities $stockpile lateral-movement
abilities $atomic lateral-movement
abilities $stockpile collection 
abilities $atomic collection
abilities $stockpile exfiltration
abilities $atomic exfiltration
abilities $stockpile impact
abilities $atomic impact
abilities $atomic multiple

# Parsing json threat model file
echo "    [!] Finding MITRE techniques in JSON model file." 

echo " " > model_parsed
cat $model_file | sed -n 's/.* "techniqueID": "\(.*\)".*/\1/p' > model_parsed
echo "      [!] Found $(cat model_parsed | wc -l) techniques."
echo "    [!] Matching CVC model techniques with Caldera abilities." 

echo " " > abilities_per_model

# Comparing CVC model to parsed abilities
for technique in $(cat model_parsed)
do
    echo $(cat $output | grep $technique | cut -d " " -f 1 ) >> abilities_per_model
done

#Creating adversary file
echo "    [!] Creating adversary file."
echo "       [?] Add custom description? [y/n]"
read custom_description
case $custom_description in
    
    "y" | "Y")
    echo "Custom Description:"
    read description
    ;;
esac

echo "---

name: $adv_name
description: $description
atomic_ordering:" > "$adv_name".yml

cat abilities_per_model | xargs -n1  | sed 's/\(.*\)/  - \1/g' >> "$adv_name".yml

echo "    [!] DONE! You file: $adv_name.yml is ready to use in Caldera."

rm model_parsed
rm abilities_parsed
rm abilities_per_model
rm platforms