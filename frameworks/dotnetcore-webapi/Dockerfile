FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS builder

WORKDIR /opt/src

COPY . /opt/src

RUN dotnet publish -c Release


FROM mcr.microsoft.com/dotnet/core/aspnet:3.1

ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080
ENV BENCH_WORKERS=1

WORKDIR /opt/benchwarmer

COPY --from=builder /opt/src/Benchwarmer/bin/Release/netcoreapp3.1/publish /opt/benchwarmer

ENTRYPOINT ["dotnet", "Benchwarmer.dll"]

