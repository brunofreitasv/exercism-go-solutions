# Stage all changes
git add .

# Commit with timestamp
$timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
git commit -m "Update Exercism solutions $timestamp"

# Push to the remote repository
git push