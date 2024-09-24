// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** Greet the user Responds with a personalized greeting message POST /api/hello */
export async function postApiHello(body: API.NameRequest, options?: { [key: string]: any }) {
  return request<API.GreetingResponse>('/api/hello', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
