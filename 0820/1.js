// const token = response.headers.get('X-Jwt-Token');
const token="eyJhbGciOiJSUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwYWFzLnBhc3Nwb3J0LmF1dGgiLCJleHAiOjE3NTU2ODU0MDUsImlhdCI6MTc1NTY4MTc0NSwidXNlcm5hbWUiOiJzdW5qaWFxaS50cGFzIiwidHlwZSI6InBlcnNvbl9hY2NvdW50IiwicmVnaW9uIjoiY24iLCJ0cnVzdGVkIjp0cnVlLCJ1dWlkIjoiY2YwZmYyYzctZDhiNS00NTVhLWE3OTUtZDFiNDY5OTU0OTQwIiwic2l0ZSI6Im9ubGluZSIsImJ5dGVjbG91ZF90ZW5hbnRfaWQiOiJieXRlZGFuY2UiLCJieXRlY2xvdWRfdGVuYW50X2lkX29yZyI6ImJ5dGVkYW5jZSIsInNjb3BlIjoiYnl0ZWRhbmNlIiwic2VxdWVuY2UiOiJUZXN0Iiwib3JnYW5pemF0aW9uIjoi5Lqn5ZOB56CU5Y-R5ZKM5bel56iL5p625p6ELeeUn-a0u-acjeWKoS3otKjph4_kv53pmpwt6LSo6YeP5p625p6EIiwid29ya19jb3VudHJ5IjoiQ0hOIiwibG9jYXRpb24iOiJDTiIsImF2YXRhcl91cmwiOiJodHRwczovL3MxLWltZmlsZS5mZWlzaHVjZG4uY29tL3N0YXRpYy1yZXNvdXJjZS92MS92M18wMGsxXzZjZDc2NTA1LWNhMGMtNGYzZi1hOGYxLWE4YjhmYWJmNjU4Z34_aW1hZ2Vfc2l6ZT1ub29wXHUwMDI2Y3V0X3R5cGU9XHUwMDI2cXVhbGl0eT1cdTAwMjZmb3JtYXQ9cG5nXHUwMDI2c3RpY2tlcl9mb3JtYXQ9LndlYnAiLCJlbWFpbCI6InN1bmppYXFpLnRwYXNAYnl0ZWRhbmNlLmNvbSIsImVtcGxveWVlX2lkIjoyNjI2MTIzLCJuZXdfZW1wbG95ZWVfaWQiOjI2MjYxMjN9.NpBri7Ju5MzdC-l9INzAH_cdApfYM8j4JFle6a_8aJMGyM5NvvfRfjUZz-jWUzxaWvL8SEDQEdskSv2U19tROunAzldnzZ95Qt1F0yVBNdb0RLBBth47r4_c7jOww0RJ0SwaJv894IkPY8QQWpHsK-oLENSTTA_xerKIKgcqYGk"

const parseJwt = ( _token_ ) => {
   try {
       const base64Url =  _token_ .split('.')[1];
       const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
       return  JSON.parse(decodeURIComponent(escape(atob(base64))));
   } catch (err) {
       console.error('jwt parse error', err);
       return  null;
   }
};
const tokenInfo = parseJwt(token);
console.log(tokenInfo);
