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
            <pre>
            {
                "to_email": string,
                "verification_link": string,
            }
            </pre>
        </td>
        <td>
            <pre>
            {
                "error": bool,
                "data": string,
            }
            </pre>
        </td>
        <td>
            Use it to send verification email to user.
            <b>Accepts request only from: <a href="#">ignitusrestapi.herokuapp.com</a> </b>
        </td>
    </tr>
</table>
