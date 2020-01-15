{{define "navbar"}}
    <a class="navbar-brand" href="/">我的新聞</a>
    <div class="collapse navbar-collapse">
    	<ul class="nav nav-pills">
        	<li class="nav-item">
        		<a {{if .IsHome}} class="nav-link active" {{else}} class="nav-link" {{end}} href="/">首頁</a>
        	</li>
        	<li class="nav-item">
        		<a {{if .IsCategory}} class="nav-link active" {{else}} class="nav-link" {{end}} href="/category">分類</a>
        	</li>
        	<li class="nav-item">
        		<a {{if .IsTopic}} class="nav-link active" {{else}} class="nav-link" {{end}} href="/topic">文章</a>
        	</li>
    	</ul>
    </div>
    
    <div class="pull-right">
    	<ul class="nav navbar-nav">
    		{{if .IsLogin}}
        		<li><a href="/login?exit=true">退出</a></li>
        	{{else}}
        		<li><a href="/login">使用者登入</a></li>
        	{{end}}
    	</ul>
    </div>
{{end}}