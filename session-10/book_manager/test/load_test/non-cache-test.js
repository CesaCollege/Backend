import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  let access_token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwaXJlZF9hdCI6MTY5MjUyOTM1OH0sInVzZXJuYW1lIjoibW9oYW1tYWQyIn0.YXrncsl00jmIgCbEgavWc7VdpiRPZsVOn_9EiOdCRkA';

  // Make an HTTP request to your target endpoint
  const headers = {
    'Authorization': access_token,
    'Use-Cache': "false"
  };
  http.get('http://localhost:8080/profile/calculate_age', { headers: headers });

  // Sleep for a short duration to simulate user behavior
  sleep(1); // Sleep for 1 second
}
