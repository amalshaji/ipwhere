<script>
	let ip = "";
	let country = "";
	let region = "";
	let isp = "";
	let message = "";
	let lat = "";
	let lon = "";

	const getData = () => {
		output.setAttribute("hidden", "");
		error.setAttribute("hidden", "");
		fetch("/api/" + ip)
			.then((res) => res.json())
			.then((data) => {
				ip = "";
				country = data.country;
				region = data.regionName;
				isp = data.isp;
				lat = data.lat;
				lon = data.lon;
				output.removeAttribute("hidden");
			})
			.catch((err) => {
				message = err.message;
				error.removeAttribute("hidden");
			});
	};
</script>

<style>
	.container-fluid {
		margin-top: 50px;
		max-height: 100%;
		margin-left: 10%;
		max-width: 80%;
	}
	#ip-add {
		max-width: 70%;
	}
</style>

<div class="container-fluid">
	<center>
		<h2>IP Geolocator</h2>
		<input type="text" bind:value={ip} id="ip-add" />
		<button type="button" class="btn btn-dark" on:click={getData}>Locate IP</button>
		<div hidden id="output">
			<div class="card">
				Country:
				<b>{country}</b>
				&nbsp; Region:
				<b>{region}</b>
				&nbsp; ISP:
				<b>{isp}</b>
			</div><br /><br />
			<iframe
				title="map"
				width="100%"
				height="200%"
				frameborder="0"
				scrolling="no"
				marginheight="0"
				marginwidth="0"
				src="https://www.openstreetmap.org/export/embed.html?bbox={lon - 0.2}%2C{lat - 0.2}%2C{lon + 0.2}%2C{lat + 0.2}&layer=mapnik&marker={lat}%2C{lon}"
				style="border: 1px solid black" /><br />
		</div>
		<div hidden id="error">{message}</div>
	</center>
</div>
