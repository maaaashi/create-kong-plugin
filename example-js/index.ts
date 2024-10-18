import kong from "kong-pdk/kong";

class KongPlugin {
  constructor(private config: any) {}
  async access(kong: kong) {
    console.log("This is an example golang plugin handler, config: " + this.config)
  }
 }
 
export default {
  Plugin: KongPlugin,
  Schema: [
    { message: { type: "string" } },
  ],
  Version: '0.1.0',
  Priority: 10,
}