using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;

namespace Benchwarmer
{
    public class Program
    {
        public static void Main(string[] args)
        {
			var host = Environment.GetEnvironmentVariable("BENCH_HOST");
			var port = Environment.GetEnvironmentVariable("BENCH_PORT");
			var addr = String.Format("http://{0}:{1}", host, port);
            CreateHostBuilder(args, addr).Build().Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args, string addr) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.UseStartup<Startup>();
					webBuilder.UseUrls(addr);
                });
    }
}
