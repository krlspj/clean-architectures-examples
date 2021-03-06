using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Serilog;
using SimpleInjector;

namespace HostApp.WebService.Client2
{
    public class Program
    {
        private static IHost _host;

        static void Main(string[] args)
        {
            ConfigureSerilog();
            Container container = new Container();

            using (_host = new ApplicationHostBuilder(args, "HostApp.WebService.Client2",  container).Build())
            {
                // Use "Start" and "WaitForShutdown" instead of "Run" as this handles being stopped
                // from a container and allow the CloseAndFlush to be called
                _host.Start();
                _host.WaitForShutdown();

                // Important to CloseAndFlush the logs inside the using to ensure all log 
                // messages from services are in the buffer before CloseAndFlush is called
                Log.Information("Shutting down...");
                Log.CloseAndFlush();
            }
        }

        private static void ConfigureSerilog()
        {
            Log.Logger = new LoggerConfiguration()
                .WriteTo.ColoredConsole()
                .MinimumLevel.Debug()
                .CreateLogger();
        }
    }
}
