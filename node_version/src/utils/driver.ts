
import { WebDriver, Capabilities, Builder } from "selenium-webdriver";

export async function buildStudentDriver(): Promise<WebDriver> {
  const caps = Capabilities.chrome();
  caps.set("nodename:applicationName", "it_admin");

  return await new Builder()
    .forBrowser("chrome")
    .withCapabilities(caps)
    .usingServer("http://localhost:4444/wd/hub")
    .build();
}
