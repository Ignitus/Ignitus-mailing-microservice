# Ignitus mailing microservice

> Scalable mailing utility.

## API

<table>
    <tr>
        <th>Endpoint</th>
        <th>Method</th>
        <th>Request Body</th>
        <th>Response</th>
        <th>Additional</th>
    </tr>
    <tr>
        <td>/mail/confirmation</td>
        <td>POST</td>
        <td>
            <code>
            {
                "to_email": string, <br/>
                "verification_link": string, <br/>
            }
            </code>
        </td>
        <td>
            <code>
            {
                "error": bool, <br/>
                "data": string, <br/>
            }
            </code>
        </td>
        <td>
            Use it to send verification email to user.
            <b>Accepts request only from: <a href="#">ignitusrestapi.herokuapp.com</a> </b>
        </td>
    </tr>
</table>

## .env

```
- SENDGRID_API_KEY
- ACCEPTED_HOST - (set localhost:8000 for local environment)
- FROM_EMAIL
```