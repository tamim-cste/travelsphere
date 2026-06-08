<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}} — TravelSphere</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    {{template "header.tpl" .}}
    
    <main>
        {{.LayoutContent}}
    </main>

    {{template "footer.tpl" .}}
</body>
</html>