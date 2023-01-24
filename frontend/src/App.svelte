<!-- <script>
  import logo from './assets/images/logo-universal.png'
  import {Bye} from '../wailsjs/go/main/App.js'

  let resultText = "Please enter your name below 👇"
  let name

  function greet() {
    Bye(name).then(result => resultText = result)
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={greet}>Bye</button>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style> -->
<script>
  import { KubeConfig, KubeSwitch } from "../wailsjs/go/main/App.js";

  import { onMount } from "svelte";

  let ctxs = []

  let configs;

  onMount(async () => {
    let configs = await KubeConfig()
    ctxs = configs.map((c,i) => ({id: i+1, text: c}))
  });

  let selected;
  let namespace; 

  let answer = "";

  async function handleSubmit() {
    console.log(selected)
    namespace = await KubeSwitch(selected.text);
  }

</script>

<h2>Switching KubeContext</h2>

<form on:submit|preventDefault={handleSubmit}>
  <select bind:value={selected} on:change={() => (answer = "")}>
    {#each ctxs as ctx}
      <option value={ctx}>
        {ctx.text}
      </option>
    {/each}
  </select>

  <!-- <input bind:value={answer}> -->

  <button disabled={false} type="submit"> Switch </button>
</form>

<p>selected: {selected ? selected.text : "[waiting...]"}</p>
<p>namespace: {namespace ? namespace : "[waiting...]"}</p>

<style>
  input {
    display: block;
    width: 500px;
    max-width: 100%;
  }
</style>
