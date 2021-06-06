package template

const INDEX = `
<!DOCTYPE html>
<html>
  <head>
    <style>
      body {
        font-family: Verdana, sans-serif;
        margin-left:50px;
        margin-right:50px;
      }

      table {
        border-collapse: collapse;
        width: 100%;
      }

      th {
        border: 1px solid #eaeaea;
        background-color: #fefefe;
        text-align: left;
        padding: 5px;
      }

      td {
        border: 1px solid #eaeaea;
        text-align: left;
        padding: 3px;
      }

      tr:nth-child(even) {
        background-color: #eaeaea;
      }

      .dir-icon {
        height: 18px;
        width: 18px;
        margin-right: 5px;
        content: url("data:image/gif;base64,R0lGODlhFAAWAMIAAP/////Mmcz//5lmMzMzMwAAAAAAAAAAACH+TlRoaXMgYXJ0IGlzIGluIHRoZSBwdWJsaWMgZG9tYWluLiBLZXZpbiBIdWdoZXMsIGtldmluaEBlaXQuY29tLCBTZXB0ZW1iZXIgMTk5NQAh+QQBAAACACwAAAAAFAAWAAADVCi63P4wyklZufjOErrvRcR9ZKYpxUB6aokGQyzHKxyO9RoTV54PPJyPBewNSUXhcWc8soJOIjTaSVJhVphWxd3CeILUbDwmgMPmtHrNIyxM8Iw7AQA7")
      }

      .file-icon {
        height: 18px;
        width: 18px;
        margin-right: 5px;
        content: url("data:image/gif;base64,R0lGODlhFAAWAMIAAP///8z//5mZmTMzMwAAAAAAAAAAAAAAACH+TlRoaXMgYXJ0IGlzIGluIHRoZSBwdWJsaWMgZG9tYWluLiBLZXZpbiBIdWdoZXMsIGtldmluaEBlaXQuY29tLCBTZXB0ZW1iZXIgMTk5NQAh+QQBAAABACwAAAAAFAAWAAADaDi6vPEwDECrnSO+aTvPEQcIAmGaIrhR5XmKgMq1LkoMN7ECrjDWp52r0iPpJJ0KjUAq7SxLE+sI+9V8vycFiM0iLb2O80s8JcfVJJTaGYrZYPNby5Ov6WolPD+XDJqAgSQ4EUCGQQEJADs=")
      }

      a:link {
        text-decoration: none;
      }
      a:visited {
        text-decoration: none;
      }

      .powered-by {
        margin-top:7px;
        font-size: 12px;
        color: gray;
      }
    </style>
  </head>
  <body>
    <h2>{{.Brand}}</h2>
    {{if eq .Dir ""}}
    <div><a href="/"><b>Parent</b></a></div>
    {{else if eq .Dir "/"}}
    <div>/</div>
    {{else}}
    <div><a href="{{.Dir}}"><b>Parent</b></a></div>
    {{end}}
    <table>
      <tr>
        <th>Name</th>
        <th>Last Modified</th>
        <th>Size</th>
      </tr>
      {{range $file := .Files}}
      <tr>
        {{ if $file.IsDir }}
        <td class="file"><span class="dir-icon"></span><a href="{{$file.FullPath}}">{{$file.Name}}</a></td>
        {{else}}
        <td class="dir"><span class="file-icon"></span><a href="{{$file.FullPath}}">{{$file.Name}}</a></td>
        {{end}}      
        <td>{{$file.ModifiedTime.Format "2006-01-02T15:04:05 -0700 MST"}}</td>
        <td>{{$file.FriendlySize}}</td>
      </tr>
      {{end}}
    </table>
    <div class="powered-by">Powered By <a href="https://github.com/jkandasa/static-file-server">Lightweight Static File Server</a></div>
  </body>
</html>
`
