<script>
  import Map from "./components/Map.svelte";
  import { emailValidator, requiredValidator } from "./lib/validators.js";
  import { createFieldValidator } from "./lib/validation.js";

  const [validity, validate] = createFieldValidator(
    requiredValidator(),
    emailValidator()
  );

  let ip = "";
  let disp_ip = "";
  let country = "";
  let region = "";
  let isp = "";
  let lat = "";
  let lon = "";
  let output;

  //   const getData = () => {
  //     output.setAttribute("hidden", "");
  //     fetch("/api/" + ip)
  //       .then((res) => res.json())
  //       .then((data) => {
  //         disp_ip = ip;
  //         country = data.country;
  //         region = data.regionName;
  //         isp = data.isp;
  //         lat = data.lat;
  //         lon = data.lon;
  //         output.removeAttribute("hidden");
  //       });
  //   };

  const getData = async () => {
    output = false;
    try {
      const res = await fetch(`/api/${ip}`);
      const data = await res.json();
      output = true;
      disp_ip = ip;
      country = data.country;
      region = data.regionName;
      isp = data.isp;
      lat = data.lat;
      lon = data.lon;
    } catch (err) {
      output = false;
      console.error(err);
    }
  };
</script>

<div class="container">
  <center>
    <h1>IP Geolocator</h1>
    <input
      type="text"
      placeholder="Enter IP address"
      bind:value={ip}
      id="ip-add"
      use:validate={ip}
    />
    <button
      disabled={!$validity.valid}
      type="button"
      class="btn btn-success"
      on:click={getData}>Locate IP</button
    >
    {#if !$validity.valid && ip.length > 0}
      <p class="error">please enter a valid ip</p>
    {/if}
    {#if output}
      <div class="card">
        IP:
        <b>{disp_ip}</b>
        Country:
        <b>{country}</b>
        &nbsp; Region:
        <b>{region}</b>
        &nbsp; ISP:
        <b>{isp}</b>
      </div>
      <br /><br />
      <Map {lat} {lon} /><br />
    {/if}
  </center>
</div>

<style>
  .container {
    margin-top: 50px;
    height: 100%;
    width: 100%;
  }
  #ip-add {
    width: 70%;
  }
  .error {
    font-size: small;
    color: red;
  }
</style>
