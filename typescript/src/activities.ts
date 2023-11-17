export async function stepOne(step: string): Promise<string> {
  console.log(`Running step ${step}!`)
  return "One done!";
}

export async function stepTwo(step: string): Promise<string> {
  console.log(`Running step ${step}!`)

  // TODO FIXME
  const hasBug = true

  if (hasBug) {
    throw new Error("‼️ ERROR!!!! Really hard to fix. 😜")
  }
  return "Two done!";
}

export async function stepThree(step: string): Promise<string> {
  console.log(`Running step ${step}!`)
  return "Three done!";
}
