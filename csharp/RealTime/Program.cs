var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Hello World!");

app.MapGet("/events", async (HttpContext ctx, CancellationToken ct) =>
{
    ctx.Response.Headers.Add("Content-Type", "text/event-stream");

    var tokens = new string [] { "this", "is", "real", "time", "response"};
    var index = 0;
    
    while (index < tokens.Length)
    {
        await ctx.Response.WriteAsync($"data: {tokens[index]}");
        await ctx.Response.WriteAsync($"\n\n");
        await ctx.Response.Body.FlushAsync();
        index++;

        //intentional sleep
        Thread.Sleep(500);
    }
});

app.Run();
