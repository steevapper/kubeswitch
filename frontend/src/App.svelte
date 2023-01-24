
<script>
  import { KubeConfig, KubeSwitch, CurrentCtx } from "../wailsjs/go/main/App.js";

  import { onMount } from "svelte";

  let ctxs = []

  let selected;
  let namespace; 
  let current;  

  onMount(async () => {
    let configs = await KubeConfig();
    ctxs = configs.map((c,i) => ({id: i+1, text: c}));
    current = await CurrentCtx();
  });

  async function handleSubmit() {
    console.log(selected)
    namespace = await KubeSwitch(selected.text);
  }

</script>

<h2>Switching KubeContext</h2>

<form on:submit|preventDefault={handleSubmit}>
  <select bind:value={selected} on:change={() => (namespace = "")}>
    {#each ctxs as ctx}
      <option value={ctx}>
        {ctx.text}
      </option>
    {/each}
  </select>

  <!-- <input bind:value={answer}> -->

  <button disabled={false} type="submit"> Switch </button>
</form>

<p>current: {current ? current : "[waiting...]"}</p>
<p>selected: {selected ? selected.text : "[waiting...]"}</p>
<p>namespace: {namespace ? namespace : "[waiting...]"}</p>

<style>
  /* input {
    display: block;
    width: 500px;
    max-width: 100%;
  } */
</style>
