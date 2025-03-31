package plugin

import (
	"context"
	"errors"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *StatefulSetRpcPlugin) lookupStatefulSet(ctx context.Context, matchLabels map[string]string, name string, namespace string) (*appsv1.StatefulSet, error) {
	r.Clienset.AppsV1().StatefulSets(namespace).Get(ctx, name, metav1.GetOptions{})
	ls := metav1.LabelSelector{
		MatchLabels: matchLabels,
	}

	labelSelector, err := metav1.LabelSelectorAsSelector(&ls)
	if err != nil {
		return nil, err
	}
	stsList, err := r.Clienset.AppsV1().StatefulSets(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		return nil, err
	}
	if len(stsList.Items) > 1 {
		return nil, errors.New("multiple StatefulSets found")
	} else if len(stsList.Items) == 0 {
		return nil, errors.New("no StatefulSet found")
	}

	ss, err := r.Clienset.AppsV1().StatefulSets(namespace).Get(ctx, stsList.Items[0].Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return ss, nil

}
