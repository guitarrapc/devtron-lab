var builder = WebApplication.CreateSlimBuilder(args);

// Add services to the container.
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();
builder.Services.AddHealthChecks();

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.TypeInfoResolverChain.Insert(0, AppJsonSerializerContext.Default);
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
}

app.MapHealthChecks("/healthz");
app.MapGet("/", () => MachineInformation.Default);
app.MapGet("/weatherforecast", () =>
{
    var forecast = Enumerable.Range(1, 5).Select(index =>
        new WeatherForecast
        (
            DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
            Random.Shared.Next(-20, 55),
            WeatherForecast.Summaries[Random.Shared.Next(WeatherForecast.Summaries.Length)]
        ))
        .ToArray();
    return forecast;
})
.WithName("GetWeatherForecast");

app.Run();

internal record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary)
{
    public static readonly string[] Summaries = ["Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"];
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}

internal record MachineInformation(string MachineName, string OSDescription, int ProcessorCount, long UsedMemoryInMB)
{
    public static MachineInformation Default { get; } = new MachineInformation(
        Environment.MachineName,
        System.Runtime.InteropServices.RuntimeInformation.OSDescription,
        Environment.ProcessorCount,
        GC.GetTotalMemory(false) / (1024 * 1024)
    );
}

[System.Text.Json.Serialization.JsonSerializable(typeof(WeatherForecast[]))]
[System.Text.Json.Serialization.JsonSerializable(typeof(MachineInformation))]
internal partial class AppJsonSerializerContext : System.Text.Json.Serialization.JsonSerializerContext
{
}
