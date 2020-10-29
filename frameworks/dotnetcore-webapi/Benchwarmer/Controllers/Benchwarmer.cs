using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace Benchwarmer.Controllers {
	[ApiController]
	[Route("/")]
	public class BenchwarmerController : ControllerBase {
		[HttpGet("/hello")]
		public string Hello() {
			return "Hello, World!";
		}

		[HttpPost("/json")]
		public JSONRes Json([FromBody] JSONReq req) {
			return new JSONRes{
				Msg = String.Format("{0} {1}", req.Greeting, req.Name)
			};
		}
	}
}
