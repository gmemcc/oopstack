#  Web development technology stack for Overseas Ops

## Backend
Any backend implementation conforms to OpenAPI 3.0 / Swagger 2.0 specification would be acceptable.
In this demo, we use Echo v4 + swaggo

## Frontend
We use Umi Max + Ant Design Pro V5 as the frontend framework. Simply run this command to generate the skeleton of the frontend project:
```bash
npm i @ant-design/pro-cli -g
pro create frontend
cd frontend
yarn
```
Now you can run the frontend project in development mode:
```bash
yarn run dev
```
To generate API stubs for the frontend project, you should first change config/config.ts to point to the backend server:
```
  openAPI: [
    {
      requestLibPath: "import { request } from '@umijs/max'",
      schemaPath: 'http://127.0.0.1:8085/swagger/doc.json',
      mock: false,
    },
  ],
```
Then run the following command:
```bash
yarn max openapi
```
You will also need to update proxy config in config/proxy.ts to point to the backend server:
```
export default {
  dev: {
    '/api/': {
      target: 'http://127.0.0.1:8085',
      changeOrigin: true,
    },
  },
};
```