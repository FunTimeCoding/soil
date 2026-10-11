import argparse
import asyncio
import re
from datetime import datetime
from pathlib import Path

from playwright.async_api import async_playwright


def parse_tid_file(path):
    content = path.read_text(encoding="utf-8")
    header, body = content.split("\n\n", 1)
    tiddler = {"type": "text/vnd.tiddlywiki", "text": body.strip()}

    for line in header.splitlines():
        match = re.match(r"^(\w+):\s*(.*)$", line)

        if match:
            key, value = match.groups()

            if key in ("title", "tags", "created", "creator"):
                tiddler[key] = value.strip()

    for key in ("title", "tags", "created", "creator"):
        if key not in tiddler:
            raise ValueError(f"missing {key} in {path.name}")

    return tiddler


async def run(args):
    template = Path(args.html).resolve().as_uri()
    log_folder = Path(args.logfolder).resolve()
    tiddlers = [parse_tid_file(p) for p in log_folder.glob("*.tid")]
    print(f"parsed {len(tiddlers)} tiddlers")

    async with async_playwright() as playwright:
        browser = await playwright.chromium.launch(headless=True)
        page = await (await browser.new_context()).new_page()
        await page.goto(template)

        for tiddler in tiddlers:
            print(f"importing {tiddler['title']}")
            await page.evaluate(
                """(tiddler) => {
                    $tw.wiki.importTiddler(new $tw.Tiddler(tiddler));
                }""",
                tiddler,
            )

        async with page.expect_download() as download_info:
            await page.click(
                "body > div.tc-page-container-wrapper > div > div > div"
                " > div.tc-topbar.tc-adjust-top-of-scroll"
                " > div.tc-topbar-left > div > button:nth-child(3)"
            )

        download = await download_info.value
        date = datetime.now().strftime("%Y-%m-%d")
        output = Path(args.download) / f"ONYX Log {date}.html"
        await download.save_as(str(output))
        print(f"saved {output}")
        await browser.close()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="TiddlyWiki tiddler importer")
    parser.add_argument("--html", required=True, help="path to TiddlyWiki template")
    parser.add_argument("--download", required=True, help="output directory")
    parser.add_argument("--logfolder", required=True, help="folder containing .tid files")
    args = parser.parse_args()
    asyncio.run(run(args))
