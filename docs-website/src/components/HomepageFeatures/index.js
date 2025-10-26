import clsx from "clsx";
import Heading from "@theme/Heading";
import styles from "./styles.module.css";

const FeatureList = [
  {
    title: "CLI-Native Learning",
    Svg: require("@site/static/img/undraw_docusaurus_mountain.svg").default,
    description: (
      <>
        Learn directly in your terminal — no browser, no cloud VMs.{" "}
        <b>TryOutShell</b> brings hands-on, guided labs where developers
        actually work.
      </>
    ),
  },
  {
    title: "Interactive & Guided",
    Svg: require("@site/static/img/undraw_docusaurus_tree.svg").default,
    description: (
      <>
        Follow step-by-step lessons that validate your commands and provide
        instant feedback — powered by a custom TUI built with Charm.sh’s Bubble
        Tea.
      </>
    ),
  },
  {
    title: "Open & Extensible",
    Svg: require("@site/static/img/undraw_docusaurus_react.svg").default,
    description: (
      <>
        All lessons and binaries are open source. Create your own YAML-based
        lessons and share them with the community.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
