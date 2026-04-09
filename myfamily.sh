curl -s "https://01.gritlab.ax/assets/superhero/all.json" | \
jq -r --arg id "$HERO_ID" '
  .[] 
  | select(.id == ($id | tonumber)) 
  | .connections.relatives 
  | if type=="array" then join("\\n")      # join array with literal \n
    elif type=="string" then gsub("\n"; "\\n")  # replace actual newlines with literal \n
    else "" 
    end
'