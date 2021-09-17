<script>
  export let lat, lon;
</script>

<iframe
				id="map"
				title="map"
				frameborder="0"
				scrolling="no"
				marginheight="0"
				marginwidth="0"
				src="https://www.openstreetmap.org/export/embed.html?bbox={lon - 0.5}%2C{lat - 0.5}%2C{lon + 0.5}%2C{lat + 0.5}&layer=mapnik▮={lat}%2C{lon}"
				style="border: 1px solid black" />

<style>
  #map {
    width: 100%;
    height: 300px;
    overflow: auto;
  }
</style>
