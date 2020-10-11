# Deploying gofiber + sveltejs on netlify

[![Netlify Status](https://api.netlify.com/api/v1/badges/143c3c42-60f7-427a-b3fd-8ca3947a2d40/deploy-status)](https://app.netlify.com/sites/gofiber-svelte/deploys)

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