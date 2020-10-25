import csv
import json
import os

reports = {}

def main():
    for fname in os.listdir("../reports"):
        parts = fname.split("_")
        framework = parts[0]
        bench = parts[1].replace(".json", "")
        with open(f"../reports/{fname}", "r") as fd:
            if not framework in reports:
                reports[framework] = {}
            reports[framework][bench] = json.load(fd)

    with open("aggregate.csv", "w") as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow([
            "framework",
            "benchmark",
            "requests",
            "duration",
            "bytes",
            "requests_per_sec",
            "bytes_per_sec",
            "latency_min",
            "latency_mean",
            "latency_max",
            ])
        for framework, benchmarks in reports.items():
            for bench, metrics in benchmarks.items():
                writer.writerow([
                    framework,
                    bench,
                    metrics["requests"],
                    metrics["duration_in_microseconds"],
                    metrics["bytes"],
                    metrics["requests_per_sec"],
                    metrics["bytes_transfer_per_sec"],
                    metrics["latency_min"],
                    metrics["latency_mean"],
                    metrics["latency_max"],
                ])


if __name__ == "__main__":
    main()
