<script>

</script>

<div>
    <a href="/">&#60;&#60;</a>
    <a href="/">&#60;</a>
    <a class="active" href="/">1</a>
    <a href="/">2</a>
    <a href="/">3</a>
    <a href="/">...</a>
    <a href="/">10</a>
    <a href="/">&#62;</a>
    <a href="/">&#62;&#62;</a>
</div>

<style>
    div {
        overflow: hidden;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
        width: 100%;
        height: 100%;
        padding: 10px;
        margin: 10px;
    }

    a {
		display: inline-block;
		color: #000;
	 	text-align: center;
	  	padding: 5px 10px;
	  	text-decoration: none;
	  	font-size: 12px;
	}

	a:hover {
	 	background-color: #ddd;
	 	color: black;
	}
	
	.active {
		color: #fff;
	 	background-color: #0a4854;
	}
</style>