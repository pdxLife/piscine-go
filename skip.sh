#ls -l | awk {print $0} 
ls -l | awk 'NR%2==1'