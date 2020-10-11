# Deploying gofiber + sveltejs on netlify

[![Netlify Status](https://api.netlify.com/api/v1/badges/f45d518a-e827-4373-9a89-53fd17879414/deploy-status)](https://app.netlify.com/sites/gofiber-svelte/deploys)

Use the template. 

Thanks to @fenny(author of gofiber) for the fiber-lambda api

### Usage

```bash
# use template
git clone 
npm run build
git add .
git commit 
git push
```

#### Netlify can build sveltejs for you, change `build.sh` to
```bash
# build.sh

npm run build
GOBIN=$(pwd)/functions go install ./...
```

I find manual build better