import { api, kv } from "@nitric/sdk";
import { connect, Profile } from "../mongo";

connect().catch((err) => console.log(err));

const helloApi = api("main");

kv("profiles").allow("get");

helloApi.get("/profile/:id", async (ctx) => {
  const { id } = ctx.req.params;

  const profile = await Profile.findById(id).exec();

  if (!profile) {
    ctx.res.status = 404;
    return ctx;
  }

  return ctx.res.json(profile);
});

helloApi.post("/profile/:id", async (ctx) => {
  const { id } = ctx.req.params;

  const { name } = ctx.req.json();

  const profile = new Profile({ _id: id, name });

  try {
    await profile.save();
  } catch (err) {
    ctx.res.status = 400;
    return ctx;
  }

  ctx.res.body = `Successfully created: ${id}`;

  return ctx;
});

helloApi.delete("/profile/:id", async (ctx) => {
  const { id } = ctx.req.params;

  await Profile.deleteOne({ _id: id }).exec();

  ctx.res.body = `Successfully deleted: ${id}`;

  return ctx;
});

helloApi.get("/profiles", async (ctx) => {
  const { prefix } = ctx.req.query;

  if (prefix) {
    const prefixReg = new RegExp(
      `^${Array.isArray(prefix) ? prefix.join("") : prefix}`
    );

    const keys = await Profile.find({ _id: { $regex: prefixReg } }).exec();

    return ctx.res.json(keys);
  }

  const keys = await Profile.find().exec();

  return ctx.res.json(keys);
});