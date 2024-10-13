import {
  bold,
  green,
  red,
  yellow,
} from "https://deno.land/std@0.113.0/fmt/colors.ts";

async function fetchDeveloperExcuse() {
  try {
    const response = await fetch("http://developerexcuses.com/");
    if (!response.ok) {
      throw new Error(`Failed to fetch: ${response.status}`);
    }
    const html = await response.text();

    // Extract excuse text from HTML using a regex pattern
    const excuseMatch = html.match(/<a.*?>(.*?)<\/a>/);
    if (excuseMatch && excuseMatch[1]) {
      const excuse = excuseMatch[1];
      console.log(green(bold("Excuse: ")) + yellow(excuse));
    } else {
      console.log(red("Unable to extract the developer excuse."));
    }
  } catch (error: unknown) {
    if (error instanceof Error) {
      console.error(red("Error fetching excuse: ") + bold(error.message));
    } else {
      console.error(
        red("Exception fetching excuse: ") +
          bold(JSON.stringify(error, null, " ")),
      );
    }
  }
}

// Run the function to fetch and print the latest developer excuse
fetchDeveloperExcuse();
