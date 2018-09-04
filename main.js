const go = new Go()
let mod, inst
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  .then(async (result) => {
    mod = result.module
    inst = result.instance
    await go.run(inst)
    inst = await WebAssembly.instantiate(mod, go.importObject)
  })
