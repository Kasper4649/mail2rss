# mail2rss
testmail to RSS, deployed on Vercel serverless

We use [testmail](https://testmail.app) and [Vercel Serverless Functions](https://vercel.com/docs/serverless-functions/introduction) for the project.

If you are a student, you can get a student plan from testmail that equals the essential plan.

Set environment variables APIKEY and NAMESPACE. See [docs](https://testmail.app/docs/#get-started) for details.

You'd better set DEPLOY_SITE as your own site.

## API
default url: https://yoursite.com/api/mail2rss to get all from testmail

params: https://yoursite.com/api/mail2rss?tag=${tag} to get specific tag mails from testmail
