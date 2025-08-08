import { buildStudentDriver } from "../utils/driver";

(async () => {
    const driver = await buildStudentDriver();
    try {
        await sleep(10000) 
        await driver.get("http://my.linewize.net/");
        const title = await driver.getTitle();
        await sleep(10000)
        console.log("Page title:", title);
    } finally {
        await driver.quit();
    }
})();

function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}